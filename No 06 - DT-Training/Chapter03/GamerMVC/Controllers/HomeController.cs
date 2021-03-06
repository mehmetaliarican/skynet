﻿using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using GamerMVC.Models; //Eklendi
using NorthwindLib; //Eklendi
using Microsoft.EntityFrameworkCore; //Eklendi (await dbSet operasyonları için)
using System.Net.Http; // Web API çağrısını yapacağımız HttpClientFactory kullanımı için eklendi
using Newtonsoft.Json; // Web API servisinden dönen JSON içeriğin Parse işlemleri için eklendi

namespace GamerMVC.Controllers
{
    /*
        HomeIndexViewModel isimli modeli kullanacak olan Controller sınıfımız
        Northwind isimli DbContext tipini kullanabilmesi için,
        Constructor üstünden enjekte ediyoruz. 
        İlgili Entity servisinin kayıt işlemi bildiğiniz üzere Startup.cs içerisinde yapılmıştı.
    */
    public class HomeController : Controller
    {
        private readonly ILogger<HomeController> _logger;
        private Northwind _db;
        private readonly IHttpClientFactory _clientFactory;
        /*
            Bu Constructor tanımı oldukça değerli. Uygulama ayağa kalkarden Startup sınıfındaki ConfigureServices metodunda 
            kayıt ettirdiğimiz nesneler Constructor Injection tekniği ile bu Controller nesnesine aktarılıyorlar.
        */
        public HomeController(ILogger<HomeController> logger, Northwind db, IHttpClientFactory clientFactory)
        {
            _logger = logger;
            _db = db;
            _clientFactory = clientFactory;
        }

        /*
            Index isimli action gerçekleştiğinde devreye giren metot.
            rModel isimli HomeIndexViewModel nesnesi tanımlayıp içerisine
            View tarafına göndermek istediğimiz model nesnelerini koyduk.

            Örneğimizde tüm firma ve oyun bilgileri döndürülüyor.

            HomeController'ın Index Action'a kadar inen View sayfası,
            View altındaki Home klasöründe yer alan Index.cshtml dosyası.

            Dolayısıyla buradaki model nesnelerini Index.cshtml'de ele alabiliriz.

            Örneği terk etmeden önce Controller metotlarını(Index, CompanyGamesDetail, CreateCompany) asenkron çağrıları kullanabilir hale getirdik.
            await ile kullanılabilir metotlar dikkat edileceği üzere ToListAsync ile bitenler.
        */
        public async Task<IActionResult> Index()
        {
            /*
                Bu kısım MVC uygulamasından, Web API çağrısının nasıl yapılabileceğini göstermek için eklenmiştir.
            */

            string uri = "api/company"; // Örnekte https://localhost:5551/api/company ile şirket bilgilerini çektiğimiz adres
            var client = _clientFactory.CreateClient(name: "GameWorldService"); // Factory üstünden HttpClient nesnesini örnekliyoruz.
            var request = new HttpRequestMessage(method: HttpMethod.Get, requestUri: uri); // HTTP Get talebimizi hazırlıyoruz
            HttpResponseMessage response = await client.SendAsync(request); // HTTP Talebini gönderiyor ve cevabını alıyoruz
            var data = await response.Content.ReadAsStringAsync(); // Gelen içeriği önce String olarak okuyoruz
            var companies = JsonConvert.DeserializeObject<IEnumerable<Company>>(data).ToList(); // ve Json formatında ters serileştirip Company türünden listeye alıyoruz

            var rModel = new HomeIndexViewModel()
            {
                Companies = companies,
                //Companies = await _db.Companies.ToListAsync(), // Web API servisi üzerinden çekildiği için bu örnekte kapatıldı.
                Games = await _db.Games.ToListAsync()
            };

            return View(rModel);
            //return View();
        }

        /*
            Home.cshtml'de firma ismine tıklanınca çalışan action metodumuz
        */
        public async Task<IActionResult> CompanyGamesDetail(int? id)
        {
            /*
            asp-route-id ile gönderilen companyID değişkeni
            varsayılan route tanımlaması gereği otomatik olarak id ismiyle içeri alınır
            */
            if (id.HasValue) //Eğer bir değere sahipse
            {
                // LINQ sorgusu ile bu firmanın oyunlarının çekelim
                var games = await (from g in _db.Games
                                   where g.CompanyID == id.Value
                                   select g).ToListAsync();

                // Hiçbir sonuç yoksa HTTP 404 NotFound dönebiliriz
                if (games.Count() == 0)
                {
                    return NotFound("Bu numaraya bağlı oyunları bulamadım");
                }
                // Eğer sonuç varsa oyun listesini View'a verebiliriz
                return View(games);
            }
            return NotFound("Bir ID girilmeli");
        }

        /*
            Yeni bir oyun firması eklerken devreye giren action metodu
        */
        [HttpPost] // Yeni bilgiler POST metodu ile gönderileceği için
        public async Task<IActionResult> CreateCompany(CompanyGameModel data)
        {
            // Eğer model veri doğrulama kuralı ihlalleri içeriyorsa
            if (!ModelState.IsValid)
            {
                // Hatalı olduğunu ve hata mesajlarını doldurup View'a döndürüyoruz.
                // View'da bu hatalar ekrana bastırılıyor.
                data.HasErrors = true;
                data.ValidationErrors = ModelState.Values.SelectMany(s => s.Errors).Select(e => e.ErrorMessage);
                return View(data);
            }

            Company company = new Company
            {
                Name = data.Name,
                Description = data.Description
            };
            _db.Companies.Add(company);
            await _db.SaveChangesAsync();

            Game game = new Game
            {
                Title = data.GameTitle,
                Year = data.GameYear,
                Popuplarity = data.GamePopularity,
                Discontinued = data.GameIsContinued,
                CompanyID = company.CompanyID
            };
            _db.Games.Add(game);
            await _db.SaveChangesAsync();

            return View(data);
        }

        /*
            CreateCompany.cshtml view'unu döndüren action metodumuz
        */
        public IActionResult CreateCompany()
        {
            return View();
        }
        public IActionResult Privacy()
        {
            return View();
        }

        [ResponseCache(Duration = 0, Location = ResponseCacheLocation.None, NoStore = true)]
        public IActionResult Error()
        {
            return View(new ErrorViewModel { RequestId = Activity.Current?.Id ?? HttpContext.TraceIdentifier });
        }
    }
}

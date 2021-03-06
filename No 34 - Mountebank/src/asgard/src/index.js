// mountebank ve kendi yazdığımız ports modülünü ekledik
const mb = require('mountebank');
const ports = require('./ports');

/* 
    Mountebank uygulamasını ayağa kaldırdığımızda, yazdığımız mock servislerin de 
    etkinleştirilmesini sağlayabiliriz.
    then fonksiyonuna odaklanın.
*/
const pingService = require('./ping-service');
const cityService = require('./city-service');

// Yeni bir mountebank örneği oluşturuyoruz
mb.create({
    port: ports.server,
    pidfile: '../mb.pid',
    logfile: '../mb.log', // Bir üst klasörde tutacağımız log dosyası bildirimi
    protofile: '../protofile.json',
    ipWhitelist: ['*']
}).then(function () {
    pingService.register(); // pingService'i 
    cityService.register(); // ve cityService'i register ediyoruz
});
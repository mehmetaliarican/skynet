from django.db import models

'''
    Burası SQLite tablolarının iz düşümü olan model sınıfılarının tutulduğu modül.
    Örneğimizde kategori ve bu kategoriye dahil olan özlü sözleri ele almaya çalışıyorum.
'''


class Category(models.Model):
    # maksimum 30 karakter uzunluğunda olabilir ve null olamaz
    title = models.CharField(max_length=30, null=False)

    # shell'deki sorgu çıktıların güzel görünsün diye __str__ fonksiyonunu eziyoruz
    def __str__(self):
        return self.title


class Quote(models.Model):
    # Quote'lar Category'ye bağlıdır. one-to-many relations söz konusu
    category = models.ForeignKey(Category, on_delete=models.CASCADE)
    # maksimum 500 karakter olabilir
    content = models.CharField(max_length=500, null=False)
    # bu da varsayılan değeri 1 olan integer bir alan
    popularity_point = models.IntegerField(default=1)
    author = models.CharField(max_length=50, null=False)

    # benzer şekilde burada da ezdik
    def __str__(self):
        return self.content

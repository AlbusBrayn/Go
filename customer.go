package main //ana fonksiyon ve en fazla 1 tane olabiliyor

import (
    "fmt" //Go dilinde ekrana yazdırabilmek ,print kullanabilmek için bu kütüphaneyi ekledim
)

type Customer struct {  // Müşteri verileri ve müşteri segmentlerini tutmak için bir yapı oluşturuyoruz. araştırmalarım sonucu Go da  Laravel Eloquent ORM gibi yapılar olduğunu fark ettim Gorm kütüphanesi
    ID        string    //string int ve. bir sürü farklı veri tipi vardır ve bunlar farklı amaçlar için kullanılır
    Age       int
    Gender    string
    Purchases int
}

type Customergorup struct { //struct ile bir veri yapısı tanımlıyoruz.
    SegmentName  string
    Customers    []Customer
    AveragePurchases float64
}

func gorupCustomersByAge(customers []Customer) []CustomerGorup {  //  önce fonksiyona isim verdim ve burada bir müşteri dizi oluşuyro bu fonksiyonda müşterileri yaşlarına göre 3 farklı gurupa ayırıyoruz
    var Gorup []CustomerGorup         //Go dilinde var ile tanımlanan değişkenler, değiştirilebiliyor const ile tanımlanan değişkenler değiştirilemiyor aynı zamanda Go dilinde, tanımlanan ancak kullanılmayan değişkenler derleyici tarafından hata olarak kabul ediliyor ve kodu çalıştırmadan direk editör üzerinden hata olduğu zaman ekranda gözüküyor
    var young CustomerGorup              
    young.GorupName = "18-25"

    var adult CustomerGorup
    adult.GorupName = "26-40"

    var senior CustomerGorup
    senior.GorupName = "41+"

    for _, customer := range customers {     //bu döngüde müşterilerin yaşlarını kontrol edip duruma göre oluşturudğum guruplardan uygun olana ekliyoruz temel bir if yapısı kullandık
        if customer.Age >= 18 && customer.Age <= 25 {
            young.Customers = append(young.Customers, customer)
        } else if customer.Age > 25 && customer.Age <= 40 {
            adult.Customers = append(adult.Customers, customer)
        } else if customer.Age > 40 {
            senior.Customers = append(senior.Customers, customer)
        }Gorup
    }

    Gorup = append(Gorup, calculateAveragePurchases(young))  //apend ksımı go dilindeki slice yapısına yeni bir eleman eklemek için kullanılıyor slice = append(slice, element)

    Gorup = append(Gorup, calculateAveragePurchases(adult))
    Gorup = append(Gorup, calculateAveragePurchases(senior))

    return gorup
}

func calculateAveragePurchases(gorup CustomerGorup) CustomerGorup {
    totalPurchases := 0 //totalPurchases adında bir değişken tanımlanır ve sıfıra eşitlenir. Bu değişken, gruptaki tüm müşterilerin alışveriş miktarlarının toplamını tutmak için kullanılır.
    for _, customer := range Gorup.Customers { //for döngüsü ile totalPurchases değişkenine müşterinin alışveriş miktarı Purchases ekliyoruz
        totalPurchases += customer.Purchases
    }
    if len(segment.Customers) > 0 {
        Gorup.AveragePurchases = float64(totalPurchases) / float64(len(gorup.Customers))
    }
    return gorup
}

func printGorupAnalysis(gorup []CustomerGorup) {
    for _, gorup := range gorup {
        fmt.Printf("Segment: %s, Average Purchases: %.2f, Number of Customers: %d\n",
            gorup.GorupName, gorup.AveragePurchases, len(gorup.Customers))
    }
}

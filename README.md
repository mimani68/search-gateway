# Search Gateway

## Specifications

* Gin framework

## Development

```bash
PORT=3000 go run main.go
```

## Usage

```bash
http :3000/ping
http :3000/product/metadata
http post :3000/search query=ترک
```

## Search types

* برند محصول را چگونه توسعه دهیم
* ظرفیت بازار محصول و سهم ما در این بازار چقدر است
* رخدادهای متناسب با محصول ما کدام است
* کجا محصول را تبلیغ کنم
* شاخص های فروش برای محصول من کدام است
* مشتریان محصول چه کسانی هستند
* رقبای داخلی و خارجی محصول چه کسانی هستند
* چه کسانی می توانند مشتریان بعدی باشند
* برای توسعه فروش محصول چه کنیم
* قیمت محصول چقدر است


## Sample dataset

```bash
http ':3001/?q=برنامه&source=zoomit&itemNumber=20'
http ':3001/?q=فناوری&source=zoomit&itemNumber=20'
http ':3001/?q=فناوری&source=farsnews&itemNumber=20'
http ':3001/?q=استارتاپ&source=tabnak&itemNumber=20'
http ':3001/?q=برنامه&source=tabnak&itemNumber=20'
http ':3001/?q=قیمت&source=tabnak&itemNumber=20'
http ':3001/?q=اقتصادی&source=tabnak&itemNumber=20'
http ':3001/?q=استارتاپ&source=iribnews&itemNumber=20'
http ':3001/?q=استارتاپ&source=irna&itemNumber=20'
http ':3001/?q=استارتاپ&source=fars&itemNumber=50'

```

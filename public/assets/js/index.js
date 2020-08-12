
$(".products1 .products1-content a.actionbg").click(function(){
    var image , name , price;
     name = $(this).parent().find('h4').first().text()
    price = $(this).parent().find('p').first().text()
    image = $(this).parent().find('img.img-responsive').first().attr('src');
    localStorage.setItem('cart' + localStorage.length++, JSON.stringify({
        image: image,
        name: name,
        price:price
    }));
    return false
});

//遍历本地存储localStorage
let carts;
let body;
let str;
for (let i = 0; i < localStorage.length; i++) {
    const key = localStorage.key(i); //获取本地存储的Key
    console.log(key);

    carts = JSON.parse(localStorage.getItem(key));
    console.log(carts);
    body += '<tr class="rem1">\n' +
        '<td class="invert-image">\n' +
        '<a href="ecommerce-single.html">\n' +
        '<img src="' +
        carts.image +
        '" alt=" " class="img-responsive" width="100px" />\n' +
        '</a>\n' +
        '</td>\n' +
        '<td class="invert product-name"><a href="ecommerce-single.html">' +
        carts.name +
        '</a> <span>In stock</span> <li>Sold by <a href="#cloudtail">Cloudtail India</a></li></td>\n' +
        '<td class="invert">\n' +
        '<div class="quantity">\n' +
        '<div class="quantity-select d-flex">\n' +
        '<div class="entry value-minus">&nbsp;</div>\n' +
        '<div class="entry value">\n' +
        '<span>1</span>\n' +
        '</div>\n' +
        '<div class="entry value-plus active">&nbsp;</div>\n' +
        '</div>\n' +
        '</div>\n' +
        '</td>\n' +
        '<td class="invert price">' +
        carts.price +
        '</td>\n' +
        '<td class="invert close">\n' +
        '<div class="rem">\n' +
        '<div class="close1" data-cartId="'+ localStorage.key(i) +'">Delete </div>\n' +
        '</div>\n' +
        '</td>\n'
}



str = '<table class="timetable_sub">\n' +
    '<thead>\n' +
    '<tr>\n' +
    '<th>Product</th>\n' +
    '<th>Product Name</th>\n' +
    '<th>Quantity</th>\n' +
    '<th>Price</th>\n' +
    '<th>Remove</th>\n' +
    '</tr>\n' +
    '</thead>\n' +
    '<tbody>\n' +
    body +
    '</tbody>\n' +
    '</table>\n';

$('.checkout-right .table-responsive').html(str);


$("td.invert.close .rem .close1").on("click", (function() {
    localStorage.removeItem($(this).data('cartid'))
    $(this).parents('tr').remove();
}));

// index.js
// 注册service worker，service worker脚本文件为sw.js

/**
 * service works area
 */
if ('serviceWorker' in navigator) {
    if (navigator.serviceWorker.controller) {
        // service worker is already registered
        console.info('active service worker is found, no need to register');
    } else {
        // register the service worker
        navigator.serviceWorker
            .register('/assets/sw.js')
            .then(reg => {
                console.log('service worker registered. With scope', reg.scope);
            })
            .catch(err => console.log('service worker not registered', err));
    }
}
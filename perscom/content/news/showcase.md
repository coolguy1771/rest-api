---
title: 'Example'
date: '2021-01-24'
author: 'Lorem Ipsum'
cover: 'static/img/hello.jpg'
description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nec interdum metus. Aenean rutrum ligula sodales ex auctor, sed tempus dui mollis. Curabitur ipsum dui, aliquet nec commodo at, tristique eget ante.'
---

## Header 2

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nec interdum metus. Aenean rutrum ligula sodales ex auctor, sed tempus dui mollis. Curabitur ipsum dui, aliquet nec commodo at, tristique eget ante. **Donec quis dolor nec nunc mollis interdum vel in purus**. Sed vitae leo scelerisque, sollicitudin elit sed, congue ante. In augue nisl, vestibulum commodo est a, tristique porttitor est. Proin laoreet iaculis ornare. Nullam ut neque quam.

> Fusce pharetra suscipit orci nec tempor. Quisque vitae sem sit amet sem mollis consequat. Sed at imperdiet lorem. Vestibulum pharetra faucibus odio, ac feugiat tellus sollicitudin at. Pellentesque varius tristique mi imperdiet dapibus. Duis orci odio, sodales lacinia venenatis sit amet, feugiat et diam.

### Header 3

Nulla libero turpis, lacinia vitae cursus ut, auctor dictum nisl. Fusce varius felis nec sem ullamcorper, at convallis nisi vestibulum. Duis risus odio, porta sit amet placerat mollis, tincidunt non mauris. Suspendisse fringilla, `odio a dignissim pharetra`, est urna sollicitudin urna, eu scelerisque magna ex vitae tellus.

```css
/* PostCSS code */

pre {
    background: #1a1a1d;
    padding: 20px;
    border-radius: 8px;
    font-size: 1rem;
    overflow: auto;

    @media (--phone) {
        white-space: pre-wrap;
        word-wrap: break-word;
    }

    code {
        background: none !important;
        color: #ccc;
        padding: 0;
        font-size: inherit;
    }
}
```

```js
// JS code

const menuTrigger = document.querySelector('.menu-trigger')
const menu = document.querySelector('.menu')
const mobileQuery = getComputedStyle(document.body).getPropertyValue(
    '--phoneWidth'
)
const isMobile = () => window.matchMedia(mobileQuery).matches
const isMobileMenu = () => {
    menuTrigger.classList.toggle('hidden', !isMobile())
    menu.classList.toggle('hidden', isMobile())
}

isMobileMenu()

menuTrigger.addEventListener('click', () => menu.classList.toggle('hidden'))

window.addEventListener('resize', isMobileMenu)
```

```html
<!-- HTML code -->

<section id="main">
    <div>
        <h1 id="title">{{ .Title }}</h1>
        {{ range .Pages }} {{ .Render "summary"}} {{ end }}
    </div>
</section>
```

```go
//Go Code
func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}


```

#### Header 4

Cillum velit officia est. Ea eiusmod elit nostrud est in labore ex. Exercitation dolor mollit dolor ad sunt aliqua quis. Officia fugiat magna exercitation non magna pariatur nostrud officia mollit tempor duis nostrud reprehenderit nisi sit. Cupidatat non irure mollit reprehenderit id aute. Aute laborum tempor consectetur enim voluptate esse veniam nostrud deserunt ut incididunt minim. Mollit fugiat proident mollit nulla amet in ullamco ullamco. Qui nulla labore enim laboris eiusmod ea excepteur. Labore quis id laborum ea aliqua exercitation irure deserunt adipisicing veniam esse laborum. Pariatur ullamco quis cupidatat eiusmod do cillum voluptate amet.

-   Aliquip culpa excepteur elit laboris non nulla aliqua culpa.
-   Exercitation consequat ullamco minim dolor ea officia.
-   Integer imperdiet turpis vitae lacus imperdiet, ut ornare ligula auctor. Integer in mi eu velit vehicula suscipit eget vulputate nulla.
-   Etiam vitae enim quis velit lobortis placerat a ut sem.
    -   Ipsum consectetur et Lorem elit.
    -   Pariatur occaecat sint voluptate.

Et ipsum pariatur fugiat culpa officia eiusmod reprehenderit voluptate ea commodo. Lorem ad magna ut minim ullamco proident aute nisi elit velit elit tempor qui. Ad ipsum cupidatat cillum sunt ad ullamco minim magna aute quis. Consectetur minim nisi qui culpa incididunt cillum in enim quis velit ipsum Lorem cupidatat reprehenderit eu. Exercitation est ipsum esse dolore labore exercitation. Pariatur elit sunt aliqua veniam enim. Laborum labore non sint reprehenderit excepteur. Duis non cillum est incididunt. Cillum id duis qui velit laboris adipisicing in ex non dolore et. Officia velit consectetur adipisicing occaecat ad aute.

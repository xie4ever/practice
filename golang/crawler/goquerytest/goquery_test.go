package goquerytest

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

// TestGoQuery1 ...
func TestGoQuery1(t *testing.T) {
	html := `<html>
			<body>
				<h1 id="title">春晓</h1>
				<p class="content1">
    			春眠不觉晓，
				处处闻啼鸟。
        		夜来风雨声，
				花落知多少。
				</p>
				<p class="content2">
    			春眠不觉晓，
				处处闻啼鸟。
        		夜来风雨声，
				花落知多少。
				</p>
			</body>
			</html>
			`
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("p:contains(春)").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(i)
		fmt.Println(selection.Text())
	})
}

// TestGoQuery2 ...
func TestGoQuery2(t *testing.T) {
	html := `
	<html><head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">

  <title>数据库内核月报 — 2014/08</title>
  <meta name="description" content="数据库内核月报，来自阿里云RDS-数据库内核组。
">

  <link rel="stylesheet" href="/monthly/css/typo.css">
  <link rel="stylesheet" href="/monthly/css/animate.css">
  <link rel="stylesheet" href="/monthly/css/main.css">
  <link rel="canonical" href="http://10.101.233.47:4000/monthly/2014/08/">
  <link rel="alternate" type="application/rss+xml" title="数据库内核月报" href="http://10.101.233.47:4000/monthly/feed.xml">

  <link rel="stylesheet" href="//cdn.staticfile.org/highlight.js/8.3/styles/tomorrow.min.css">
  <script async="" src="//www.google-analytics.com/analytics.js"></script><script src="/monthly/js/highlight.min.js"></script>
  <!-- <link rel="stylesheet" href="/monthly/themes/tomorrow.css">
  <script src="/monthly/highlight/highlight.pack.js"> -->
  <script>hljs.initHighlightingOnLoad();</script>

  <script src="http://cdn.staticfile.org/jquery/1.11.1/jquery.min.js"></script>
  <script src="http://cdn.staticfile.org/jquery/1.11.1/jquery.min.map"></script>

  <script src="/monthly/scripts/changeTarget.js"></script>
  
<script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

  ga('create', 'UA-62056244-1', 'auto');
  ga('send', 'pageview');
</script></head>


<!-- Google Analysis -->



  <body>

    <header>

  <a id="go-back-home" href="/monthly/">
    <h1>数据库内核月报</h1>
  </a>

</header>


        <div id="container" class="animated zoomIn">
  <div class="block">
    <h1>
      数据库内核月报 － 2014 / 08
    </h1>
  </div>

  <div class="content typo">

    <ul class="posts">
      
<li>
  <h3>
    <small class="datetime muted" data-time="2014-08-01 00:00:00 +0800"># 01 </small>
    <a href="/monthly/2014/08/01/" target="_blank">
      
      MySQL · 参数故事 · timed_mutexes
    </a>
  </h3>
</li>

<li>
  <h3>
    <small class="datetime muted" data-time="2014-08-02 00:00:00 +0800"># 02 </small>
    <a href="/monthly/2014/08/02/" target="_blank">
      
      MySQL · 参数故事 · innodb_flush_log_at_trx_commit
    </a>
  </h3>
</li>

<li>
  <h3>
    <small class="datetime muted" data-time="2014-08-03 00:00:00 +0800"># 03 </small>
    <a href="/monthly/2014/08/03/" target="_blank">
      
      MySQL · 捉虫动态 · Count(Distinct) ERROR
    </a>
  </h3>
</li>

<li>
  <h3>
    <small class="datetime muted" data-time="2014-08-04 00:00:00 +0800"># 04 </small>
    <a href="/monthly/2014/08/04/" target="_blank">
      
      MySQL · 捉虫动态 · mysqldump BUFFER OVERFLOW
    </a>
  </h3>
</li>

<li>
  <h3>
    <small class="datetime muted" data-time="2014-08-05 00:00:00 +0800"># 05 </small>
    <a href="/monthly/2014/08/05/" target="_blank">
      
      MySQL · 捉虫动态 · long semaphore waits
    </a>
  </h3>
</li>

<li>
  <h3>
    <small class="datetime muted" data-time="2014-08-06 00:00:00 +0800"># 06 </small>
    <a href="/monthly/2014/08/06/" target="_blank">
      
      MariaDB · 分支特性 · 支持大于16K的InnoDB Page Size
    </a>
  </h3>
</li>

<li>
  <h3>
    <small class="datetime muted" data-time="2014-08-07 00:00:00 +0800"># 07 </small>
    <a href="/monthly/2014/08/07/" target="_blank">
      
      MariaDB · 分支特性 · FusionIO特性支持
    </a>
  </h3>
</li>

<li>
  <h3>
    <small class="datetime muted" data-time="2014-08-08 00:00:00 +0800"># 08 </small>
    <a href="/monthly/2014/08/08/" target="_blank">
      
      TokuDB · 性能优化 · Bulk Fetch
    </a>
  </h3>
</li>

<li>
  <h3>
    <small class="datetime muted" data-time="2014-08-09 00:00:00 +0800"># 09 </small>
    <a href="/monthly/2014/08/09/" target="_blank">
      
      TokuDB · 数据结构 · Fractal-Trees与LSM-Trees对比
    </a>
  </h3>
</li>

<li>
  <h3>
    <small class="datetime muted" data-time="2014-08-10 00:00:00 +0800"># 10 </small>
    <a href="/monthly/2014/08/10/" target="_blank">
      
      TokuDB·社区八卦·TokuDB团队
    </a>
  </h3>
</li>


    </ul>

  </div>
<!--
  <div class="block">

  </div>
-->
</div>


    <footer>
  <script src="https://cdn1.lncld.net/static/js/av-mini-0.6.10.js"></script>
  <script src="http://jerry-cdn.b0.upaiyun.com/hit-kounter/hit-kounter-lc-0.2.0.js"></script>
  <a href="http://mysql.taobao.org/" target="_blank" class="muted">阿里云RDS-数据库内核组</a>
  <br>
  <a href="https://github.com/alibaba/AliSQL" target="_blank" class="muted">欢迎在github上star AliSQL</a>
  <br>
  阅读：<span data-hk-page="current"> - </span><span class="pause">  </span>
<br>
<a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/3.0/" target="_blank"><img alt="知识共享许可协议" style="border-width:0" src="https://i.creativecommons.org/l/by-nc-sa/3.0/88x31.png"></a><br>本作品采用<a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/3.0/" target="_blank">知识共享署名-非商业性使用-相同方式共享 3.0 未本地化版本许可协议</a>进行许可。
</footer>

<script type="text/javascript">
  jQuery(document).ready(function($){
    // browser window scroll (in pixels) after which the "back to top" link is shown
    var offset = 300,
      //browser window scroll (in pixels) after which the "back to top" link opacity is reduced
      offset_opacity = 1200,
      //duration of the top scrolling animation (in ms)
      scroll_top_duration = 700,
      //grab the "back to top" link
      $back_to_top = $('.cd-top');

    //hide or show the "back to top" link
    $(window).scroll(function(){
      ( $(this).scrollTop() > offset ) ? $back_to_top.addClass('cd-is-visible') : $back_to_top.removeClass('cd-is-visible cd-fade-out');
      if( $(this).scrollTop() > offset_opacity ) {
        $back_to_top.addClass('cd-fade-out');
      }
    });

    //smooth scroll to top
    $back_to_top.on('click', function(event){
      event.preventDefault();
      $('body,html').animate({
        scrollTop: 0 ,
        }, scroll_top_duration
      );
    });

  });
</script>



    <a href="#0" class="cd-top"><svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="10px" width="38px" height="60px" viewBox="0 0 16 16" enable-background="new 0 0 16 16" xml:space="preserve">
      <polygon fill="#FFFFFF" points="8,2.8 16,10.7 13.6,13.1 8.1,7.6 2.5,13.2 0,10.7 "></polygon>
    </svg>
    </a>
  


</body></html>
`
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("a:contains(·)").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		text = getTitle(text)
		fmt.Println(text)
	})
}

func getTitle(str string) string {
	str = strings.Replace(str, "\n", "", -1)
	rs := []rune(str)

	var (
		firstIdx     int
		lastIdx      int
		count        int
		lastPointIdx int
	)

	for idx, r := range rs {
		if firstIdx == 0 && r != ' ' {
			firstIdx = idx
		}

		if r != ' ' {
			lastIdx = idx
		}

		if r == '·' {
			count++
			lastPointIdx = idx
		}
	}

	prefix := string(rs[firstIdx : lastPointIdx+2])
	prefix = strings.Replace(prefix, " ", "", -1)
	suffix := string(rs[lastPointIdx+2 : lastIdx+1])

	return fmt.Sprintf("%s%s", prefix, suffix)
}

// TestGoQuery3 ...
func TestGoQuery3(t *testing.T) {
	url := "http://mysql.taobao.org/monthly/2019/03/"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	dom, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("a:contains(·)").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		text := selection.Text()
		text = getTitle(text)
		fmt.Println(fmt.Sprintf("%s%s %s", "http://mysql.taobao.org/", href, text))
	})
}

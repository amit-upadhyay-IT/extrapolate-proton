package utils

func InHTMLTag(text string) string {
	return `
<HTML>
<style>
#p1 {background-color:rgb(255,0,0);opacity:0.6;}
#p2 {background-color:rgb(0,255,0);opacity:0.6;}
#p3 {background-color:rgb(0,0,255);opacity:0.6;}
#p4 {background-color:rgb(192,192,192);opacity:0.6;}
#p5 {background-color:rgb(255,255,0);opacity:0.6;}
#p6 {background-color:rgb(255,0,255);opacity:0.6;}
#button_text {font-size: 25px;}

.collapsible {
  background-color: #777;
  color: white;
  cursor: pointer;
  padding: 18px;
  width: 100%;
  border: none;
  text-align: left;
  outline: none;
  font-size: 15px;
}

.active, .collapsible:hover {
  background-color: #555;
}

.content {
  padding: 0 18px;
  display: none;
  overflow: hidden;
  background-color: #f1f1f1;
}

.contentwithoutpadding {
  padding: 0 0px;
  display: none;
  overflow: hidden;
  background-color: #f1f1f1;
}

.notebox {
  padding: 10px;
  border: 5px solid gray;
  margin: 0;
}
</style>
`+
		text+ `
<script>
var coll = document.getElementsByClassName("collapsible");
var i;

for (i = 0; i < coll.length; i++) {
  coll[i].addEventListener("click", function() {
    this.classList.toggle("active");
    var content = this.nextElementSibling;
    if (content.style.display === "block") {
      content.style.display = "none";
    } else {
      content.style.display = "block";
    }
  });
}
</script>
</HTML>
`
}

func QuestionHeading(text interface{}) string {
	return "<h2>" + text.(string) + "</h2>"
}

func NormalBreak() string {
	return "</br>"
}

func DifferentSolTypeBreak() string {
	return "<hr style=\"height:2px;border-width:0;color:gray;background-color:gray\">"
}

func AfterQuestionBreak() string {
	return "<hr style=\"height:5px;border-width:0;color:gray;background-color:gray\">"
}

func CodeFormattingWithBoundary(text string) string {
	return "<div class=\"notebox\"><pre><code>" + text + "</code></pre><br></div>"
}

func CodeFormatting(text interface{}) string {
	return "<pre><code>" + text.(string) + "</code></pre>"
}


func InsideOutputBox(text interface{}) string {
	return "<h3 id=\"p5\">" + text.(string) + "</h3></br>"
}

func WrapH3Text(text interface{}) string {
	return "<h3>" + text.(string) + "</h3>"
}

func WrapH4Text(text interface{}) string {
	return "<h4>" + text.(string) + "</h4>"
}

func Collapsible(heading string, text interface{}, includeBreak bool) string {
	var res string
	if includeBreak {
		res += "</br>"
	}
	res += "<button type=\"button\" class=\"collapsible\" id=\"button_text\">"+heading+"</button>" + "<div class=\"content\">"+text.(string)+"</div>"
	return res
}

func CollapsibleWithoutPadding(heading string, text interface{}, includeBreak bool) string {
	var res string
	if includeBreak {
		res += "</br>"
	}
	res += "<button type=\"button\" class=\"collapsible\" id=\"button_text\">"+heading+"</button>" + "<div class=\"contentwithoutpadding\">"+text.(string)+"</div>"
	return res
}

func NoteBox(text string) string {
	return "</br>" + "<pre><div class=\"notebox\">" + text + "</div></pre>"
}

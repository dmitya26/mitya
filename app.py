from flask import Flask, render_template
import json

app = Flask(__name__)

@app.route("/")
def index():
    return render_template("index.html")

@app.route("/about")
def about():
    return render_template("about.html")

@app.route("/writing/<titleSeperatedWithHyphens>")
def writings(titleSeperatedWithHyphens):
    title = " ".join(titleSeperatedWithHyphens.split("-"))
    date=""
    body=""

    with open("/root/mitya/book.json") as f:
        data = json.load(f)
        date=data[titleSeperatedWithHyphens]["date"]
        body=data[titleSeperatedWithHyphens]["body"]

    return render_template("individualWriting.html", date=date, title=title, body=body)

@app.route("/all-writings")
def allWritings():
    bookData = []
    with open("/root/mitya/book.json") as f:
        data = json.load(f)
        for k in data.keys():
            title = " ".join(k.split("-"))
            bookData.append((title, f"/{k}"))

    return render_template("displayWritings.html", bookData=bookData)

@app.route("/reading")
def reading():
    return render_template("reading.html")

@app.route("/portfolio")
def portfolio():
    return render_template("portfolio.html")

@app.route("/contact")
def contact():
    return render_template("contact.html")

if __name__ == "__main__":
    app.run(debug=False, port="8080", host="localhost")

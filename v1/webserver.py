from flask import Flask, render_template, request, jsonify
from v1.madmapper import send_theme

app = Flask(__name__)

@app.route("/", methods=['GET'])
def m():
    return render_template('index.html')

@app.route("/buttons", methods=['GET'])
def but():
    return render_template('buttons.html')

@app.route("/themes", methods=['POST'])
def but_post():
    content = request.get_json()['theme']
    send_theme(content)

    return "OK"

app.run(host='0.0.0.0', port=5000)
# app.run(debug = True)

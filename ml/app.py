from flask import Flask, jsonify, request
from flask_cors import CORS
import bert

app = Flask(__name__)
CORS(app)


@app.post("/api/tokenize")
def predict():
    jsonData = request.json if request.json is not None else {}
    sentence = jsonData["sentence"]
    sentence = bert.remove_control_words(sentence)
    results = bert.find_symptoms(sentence)
    print(results)

    return jsonify({"tokens": results})


if __name__ == "__main__":
    app.run(debug=True)

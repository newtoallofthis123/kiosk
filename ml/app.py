from flask import Flask, jsonify, request
from flask_cors import CORS
import utils

app = Flask(__name__)
CORS(app)


@app.post("/api/tokenize")
def predict():
    sentence = request.json["sentence"]
    sentence = utils.remove_control_words(sentence)
    results = utils.find_symptoms(sentence)
    print(results)

    return jsonify({"tokens": results})


if __name__ == "__main__":
    app.run(debug=True)

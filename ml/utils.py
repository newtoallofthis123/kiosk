from fuzzywuzzy import process
import nltk
from nltk.corpus import stopwords


def get_symptoms():
    with open(".words", "r") as f:
        symptoms = f.readlines()

    return [symptom.strip() for symptom in symptoms]


def find_symptoms(sentence):
    matched_symptoms = set()
    symptoms = get_symptoms()

    matched = process.extract(sentence, symptoms)
    for symptom, score in matched:
        if score > 80:
            matched_symptoms.add(symptom)

    return list(matched_symptoms)


# nltk.download("stopwords")


def remove_control_words(sentence):
    stop_words = set(stopwords.words("english"))

    words = sentence.split()

    filtered_words = [word for word in words if word.lower() not in stop_words]

    filtered_sentence = " ".join(filtered_words)

    return filtered_sentence

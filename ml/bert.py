from fuzzywuzzy import process
import nltk
from nltk.corpus import stopwords
import pandas as pd

def get_general_symptoms():
    with open(".words", "r") as f:
        symptoms = f.readlines()

    return [symptom.strip() for symptom in symptoms]

def get_symptoms():
    df = pd.read_csv("Training.csv")
    # get all of the column names
    symptoms = list(df.columns[:-1])
    symptoms = [symptom.replace("_", " ") for symptom in symptoms]
    symptoms.extend(get_general_symptoms())
    return symptoms


def find_symptoms(sentence):
    matched_symptoms = set()
    symptoms = get_symptoms()

    matched = process.extract(sentence, symptoms)
    for symptom, score in matched: # type: ignore
        if score > 90:
            matched_symptoms.add(symptom)

    return list(matched_symptoms)


# nltk.download("stopwords")


def remove_control_words(sentence):
    stop_words = set(stopwords.words("english"))

    words = sentence.split()

    filtered_words = [word for word in words if word.lower() not in stop_words]

    filtered_sentence = " ".join(filtered_words)

    return filtered_sentence

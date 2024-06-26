import pandas as pd
import numpy as np
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import accuracy_score, classification_report
import joblib


def read_data():
    df = pd.read_csv("./Training.csv", sep=",", header=0)
    return df


def main():
    df = read_data()
    X = df.drop(columns=["prognosis", "Unnamed: 133"])
    y = df["prognosis"]

    from sklearn.preprocessing import LabelEncoder

    le = LabelEncoder()
    y = le.fit_transform(y)

    X_train, X_test, y_train, y_test = train_test_split(
        X, y, test_size=0.2, random_state=42
    )

    clf = RandomForestClassifier(n_estimators=100, random_state=42)
    clf.fit(X_train, y_train)

    y_pred = clf.predict(X_test)

    accuracy = accuracy_score(y_test, y_pred)
    print(f"Accuracy: {accuracy:.2f}")
    print("Classification Report:")
    print(classification_report(y_test, y_pred, target_names=le.classes_))

    joblib.dump(clf, "disease_prediction_model.pkl")
    joblib.dump(le, "label_encoder.pkl")

    return


def predict_disease(symptom_names):
    all_symptoms = list(pd.read_csv("./Training.csv", sep=",", header=0).columns[:-2])
    clf = joblib.load("disease_prediction_model.pkl")
    le = joblib.load("label_encoder.pkl")

    symptoms_input = np.zeros(len(all_symptoms))

    for symptom in symptom_names:
        if symptom in all_symptoms:
            index = all_symptoms.index(symptom)
            symptoms_input[index] = 1
        else:
            print(f"Warning: Symptom '{symptom}' not recognized.")

    # symptoms_input = symptoms_input.reshape(1, -1)
    symptoms_df = pd.DataFrame([symptoms_input], columns=all_symptoms)

    prediction = clf.predict(symptoms_df)

    predicted_disease = le.inverse_transform(prediction)

    return predicted_disease


if __name__ == "__main__":
    symptoms = ["cough", "watering_from_eyes"]
    predicted_disease = predict_disease(symptoms)
    print(f"The predicted disease is: {predicted_disease}")
    # main()


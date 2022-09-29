from transformers import pipeline
classifier = pipeline("text-classification",model='bhadresh-savani/bert-base-uncased-emotion', top_k=1)
prediction = classifier("Good Good Study, Day Day Up", )
print(prediction)

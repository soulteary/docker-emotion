from transformers import pipeline
translator = pipeline("translation", model="penpen/novel-zh-en", max_time=7)
prediction = translator("张无忌抄起一张板凳，将成昆拍晕了过去。", )[0]["translation_text"]
print(prediction)

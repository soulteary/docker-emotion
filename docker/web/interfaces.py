import gradio as gr
from transformers import pipeline

classifier = pipeline("text-classification", model="bhadresh-savani/bert-base-uncased-emotion", top_k=1)
translator = pipeline("translation", model="penpen/novel-zh-en", max_time=7)

def analytics_emo(x):
    data = classifier(x)[0]
    return data

def predict(text):
    translation = ""
    split_text = text.splitlines()
    for text in split_text:
        text = text.strip()
        if text:
            if len(text) < 512:
                sentence = translator(text)[0]["translation_text"] + "\n\n"
                translation += sentence
                print(split_text)
            else:
                for i in range(0, len(text), 512):
                    if i + 512 > len(text):
                        sentence = translator(text[i:])[0]["translation_text"]
                    else:
                        sentence = translator(
                            text[i: i + 512])[0]["translation_text"]
                    translation += sentence
    return translation


with gr.Blocks() as demo:
    gr.Markdown("<center><h1>内容情感分析</h1> 一个简单的文本情感分析工具</center>")
    with gr.Tab("情感分析"):
        with gr.Row():
            with gr.Column(scale=1, min_width=600):
                translate_input = gr.Textbox(label="文本内容", lines=4, max_lines=100, placeholder="等待分析的文本内容...")
                translate_button = gr.Button("翻译")
                translate_hidden = gr.State("")
                translate_output = gr.Textbox(label="分析结果", lines=4, max_lines=100, placeholder="分析结果...")
                analytics_button = gr.Button("一窥究竟")
            text_output = gr.Textbox(label="结果", lines=10, max_lines=100, placeholder="分析结果...")
    translate_button.click(predict, api_name="translate", inputs=[translate_input], outputs=translate_output)
    analytics_button.click(analytics_emo, api_name="analytics", inputs=translate_output, outputs=text_output)

demo.launch(debug=True, server_name="0.0.0.0")

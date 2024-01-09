import os

ruta_actual = os.getcwd()
# print(ruta_actual)

cnt = 0
reject = 0
names = [
    "Message-ID:",
    "Date:",
    "From:",
    "To:",
    "Subject:",
    "Mime-Version:",
    "Content-Type:",
    "Content-Transfer-Encoding:",
    "X-From:",
    "X-To:",
    "X-cc:",
    "X-bcc:",
    "X-Folder:",
    "X-Origin:",
    "X-Filename:"
]
def navigate_file(file):
    global cnt, reject
    cnt += 1
    with open(file, 'r', encoding='utf-8', errors='replace') as archivo:
        lineas = archivo.readlines()
        for i in range(15):
            if lineas[i][:3] == "Cc:":
                print(lineas[i], names[i], lineas)
                print(i)
                assert()
        # print(lineas[0])

def get_files(ruta):
    # print(ruta)
    for file in os.listdir(ruta):
        ruta_completa = os.path.join(ruta, file)
        # print(ruta_completa)

        if os.path.isfile(ruta_completa):
            # print("is file" , ruta_completa)
            navigate_file(ruta_completa)
        elif os.path.isdir(ruta_completa):
            # print("is dir" , ruta_completa)
            get_files(ruta_completa)

get_files(os.path.join(ruta_actual, "enron_mail_20110402"))
print(cnt)
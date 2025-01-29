import tkinter as tk
from tkinter import filedialog
import sys

def browse_file():
    filename = filedialog.askopenfilename(
        initialdir="/",
        title="Выберите файл",
        filetypes=(("Текстовые файлы", "*.txt"), ("все файлы", "*.*"))
    )
    if filename:
        entry.delete(0, tk.END)
        entry.insert(0, filename)

def ok_pressed():
    global root  # Объявляем root как глобальную
    filename = entry.get()
    if filename:
        print(filename, end="")
        sys.stdout.flush()
        root.destroy()  

root = tk.Tk()
root.title("Простое окно")
root.geometry("300x200")

entry = tk.Entry(root, width=50)
entry.pack()

browse_button = tk.Button(root, text="Обзор", command=browse_file)
browse_button.pack()

ok_button = tk.Button(root, text="OK", command=ok_pressed)
ok_button.pack()

root.mainloop() 
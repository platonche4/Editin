import tkinter as tk
from tkinter import scrolledtext, Menu
import webbrowser

class TextEditor:
    def __init__(self, root):
        self.root = root
        self.root.title("Текстовый редактор")
        self.root.geometry("600x400")

        # Настройка темной темы
        self.root.configure(bg='black')

        self.text_area = scrolledtext.ScrolledText(self.root, bg='black', fg='white', 
                                                    insertbackground='white', wrap='word')
        self.text_area.pack(expand=True, fill='both')

        # Добавление меню
        menu = Menu(self.root)
        self.root.config(menu=menu)

        file_menu = Menu(menu)
        menu.add_cascade(label="Файл", menu=file_menu)
        file_menu.add_command(label="Выход", command=self.root.quit)

        help_menu = Menu(menu)
        menu.add_cascade(label="Справка", menu=help_menu)
        help_menu.add_command(label="Документация", command=self.open_documentation)

    def open_documentation(self):
        webbrowser.open("https://www.example.com/documentation")  # Замените на нужный URL

if __name__ == "__main__":
    root = tk.Tk()
    editor = TextEditor(root)
    root.mainloop()
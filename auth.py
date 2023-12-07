default_username = "random"
default_password = "strongpassword"

def auth():
    username = input(f"Veuillez entrer votre nom: ")
    user_password = input(f"Veuillez confirmer votre identité en tapant le mot de passe: ")

    if username != default_username or user_password != default_password:
        print(f"Oups, vous n'avez pas acces à ce wallet, sorry. ")
        return False
    else:
        return True
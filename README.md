Imaginons un supermarché.

Step 1

Le supermarché a une liste de produits chaque un avec son price et quantité en stock.

Un acheteur arrive au supermarché et achete quelques produits en stock.

creer un petit programme qui permette au supermarche de generer le total acheté et de retirer du stock.

But: decouvrir types de donées, operations basiques et notation

Step 2

Le supermarché a ouvert son outlet en exclusivité sur certain produits.
En plus de la promotion pour chaque produit, il y a aussi une policie de fidelité:
Si vous achetez au moins 2 types de produit en promotion vous pouvez avoir $5 de réduction en plus.
Si vous achetez au moins 3  types de produit, vous aurez $5, et 10% de réduction sur le total en plus.
Si  vous achetez au moins 4  types de produit, vous aurez $5 et $2 pour chaque type different de produit en promotion, en plus.

But: découvrir les variadic functions et les multiple return values et le switch

Step 3
Nous allons publier une interface rest basique pour gerer notre cadie.
Nous allons utiliser dep pour creer notre projet
L'interface aura une methode post pour acheter un produit une methode get pour obtenir la liste de produits disponibles
Nous allons utiliser la dependence "github.com/gorilla/mux" pour créer le router

hint: https://github.com/golang/dep

 1. dans le repertoire principal du projet:

````
    $ dep init
````
ça fait la creation de plusieurs fichiers et un repertoire "vendor".

La commande principale à utiliser:

````
    $ dep ensure
````
vous pouvez la tester:

````
    $ dep help ensure
````

et vous pouvez regarder l'état des dépendences

````
    $ dep status
````

vous pouvez installer graphviz
````
    $ brew install graphviz
````

et visualizer les dépendences

````
    $ dep status -dot | dot -T png | open -f -a /Applications/Preview.app
````



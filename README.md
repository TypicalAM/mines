# 💣 Mines

👋 Hi! This is a small mines game written in golang. It uses a go port for a game library called [raylib-go](https://github.com/gen2brain/raylib-go). I created it to try out game development and **learn a new language** in the process! I also remade some of the features from [raygui](https://github.com/raysan5/raygui) to make the creating GUI widgets much **easier** and more **customizable**.

## 🧐 What are the features?

This mines game allows you to:

- Play mines! (obviously)
- Create your own board configurations
- Use and create **custom themes**!
- Make it to the scoreboard and try to get the best time on your favorite mines board

## ✨ Downloading

### Releases

You can download a fresh copy of `mines` on our [releases](https://github.com/TypicalAM/mines/releases) page. The game is built both on windows and on linux. Open the executable and you're good to *go*!

### Development

Clone the repository:

```
git clone https://github.com/TypicalAM/mines && cd mines
```

Compile for linux:

```
go run main.go
```

If the compilation fails, then install the dependencies for [raylib-go](https://github.com/gen2brain/raylib-go) because to include the library you have to generate the `cgo` definitions for the raylib C library. For cross compiling to windows from linux you can use the following command

```
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=/bin/x86_64-w64-mingw32-gcc go build -o mines.exe main.go
```

## 📸 Product images

<p align="center">
    <img src="assets/title.png" />
</p>
<p align="center">
    <img src="assets/gameplay.png" />
</p>
<p align="center">
    <img src="assets/options.png" />
</p>
<p align="center">
    <img src="assets/lost.png" />
</p>
<p align="center">
    <img src="assets/leaderboards.png" />
</p>

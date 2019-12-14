<template>
  <div id="app">
    <router-view />
  </div>
</template>

<style lang="scss">
.orange {
  color: #e75138;
}

@font-face {
  font-family: "Circe";
  src: url("./assets/fonts/Circe-Regular.eot");
  src: local("Circe"), local("Circe-Regular"),
    url("./assets/fonts/Circe-Regular.eot?#iefix") format("embedded-opentype"),
    url("./assets/fonts/Circe-Regular.woff") format("woff"),
    url("./assets/fonts/Circe-Regular.ttf") format("truetype");
  font-weight: normal;
  font-style: normal;
}

@font-face {
  font-family: "Circe";
  src: url("./assets/fonts/Circe-ExtraLight.eot");
  src: local("Circe ExtraLight"), local("Circe-ExtraLight"),
    url("./assets/fonts/Circe-ExtraLight.eot?#iefix")
      format("embedded-opentype"),
    url("./assets/fonts/Circe-ExtraLight.woff") format("woff"),
    url("./assets/fonts/Circe-ExtraLight.ttf") format("truetype");
  font-weight: 200;
  font-style: normal;
}

@font-face {
  font-family: "Circe";
  src: url("./assets/fonts/Circe-Thin.eot");
  src: local("Circe Thin"), local("Circe-Thin"),
    url("./assets/fonts/Circe-Thin.eot?#iefix") format("embedded-opentype"),
    url("./assets/fonts/Circe-Thin.woff") format("woff"),
    url("./assets/fonts/Circe-Thin.ttf") format("truetype");
  font-weight: 100;
  font-style: normal;
}

@font-face {
  font-family: "Circe";
  src: url("./assets/fonts/Circe-Light.eot");
  src: local("Circe Light"), local("Circe-Light"),
    url("./assets/fonts/Circe-Light.eot?#iefix") format("embedded-opentype"),
    url("./assets/fonts/Circe-Light.woff") format("woff"),
    url("./assets/fonts/Circe-Light.ttf") format("truetype");
  font-weight: 300;
  font-style: normal;
}

@font-face {
  font-family: "Circe";
  src: url("./assets/fonts/Circe-Bold.eot");
  src: local("Circe Bold"), local("Circe-Bold"),
    url("./assets/fonts/Circe-Bold.eot?#iefix") format("embedded-opentype"),
    url("./assets/fonts/Circe-Bold.woff") format("woff"),
    url("./assets/fonts/Circe-Bold.ttf") format("truetype");
  font-weight: bold;
  font-style: normal;
}

@font-face {
  font-family: "Circe";
  src: url("./assets/fonts/Circe-ExtraBold.eot");
  src: local("Circe ExtraBold"), local("Circe-ExtraBold"),
    url("./assets/fonts/Circe-ExtraBold.eot?#iefix") format("embedded-opentype"),
    url("./assets/fonts/Circe-ExtraBold.woff") format("woff"),
    url("./assets/fonts/Circe-ExtraBold.ttf") format("truetype");
  font-weight: 800;
  font-style: normal;
}

#app {
  font-family: Circe;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>

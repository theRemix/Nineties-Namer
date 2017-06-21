<template>
  <div class="about">
    <div class="inner">
      <h1>
        <span class="pink">90s</span>
        <span class="blue">Namer</span>
      </h1>
      <h3>Generate random names from the 90s! <span>as a service™</span></h3>
      <h3>Need some randomly generated names? 90s Namer will generate random names based on a hash or any input.</h3>
      <p class="contribute button">
        <a href="https://github.com/theRemix/Nineties-Namer">Contribute</a>
      </p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'about'
}
</script>

<style scoped>
h1 {
  border-bottom: 5px solid black;
  text-align: center;
  text-transform: uppercase;
  margin: 0;
  font-weight: normal;
  margin-top: -30px;
}
h1 span {
  display: block;
}
.pink {
  color: #E22395;
  font-family: 'Monoton', sans-serif;
  margin-bottom: -30px;
  font-size: 150px;
}
.blue {
  color: #1AD6C7;
  font-family: 'Raleway', sans-serif;
  font-size: 60px;
  letter-spacing: 5px;
  position: relative;
  bottom: 60px;
  left: 50px;
  font-weight: bold;
  margin-bottom: -50px;
}
h3 span {
  font-style: italic;
}
.about {
  font-size: 12px;
  font-family: 'Raleway', sans-serif;
}
.about .inner {
  padding: 0;
}
@media (min-width: 930px) {
  .about {
    width: 30%;
    margin-right: 20px;
  }
  .about .inner {
    padding: 20px;
  }
  .pink {
    font-size: 100px;
  }
  .blue {
    font-weight: 200;
    margin-bottom: 10px;
    bottom: 0;
    left: 0;
  }
}
</style>

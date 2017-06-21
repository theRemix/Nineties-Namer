<template>
  <div class="about">
    <div class="inner">
      <h1>
        <span class="pink">90s</span>
        <span class="blue">Namer</span>
      </h1>
      <h3>Generate random names from the 90s! <span>as a service™</span></h3>
      <h3>Need some randomly generated names? 90s Namer will generate random names based on a hash or any input.</h3>
      <p>
        <a class="contribute button" href="https://github.com/theRemix/Nineties-Namer">Contribute</a>
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
  margin: 0;
  font-weight: normal;
  margin-top: -40px;

  padding-bottom: 20px;
  margin-bottom: 25px;
}

@media (min-width: 930px) {
  h1{
    text-align: left;
    margin-top: -20px;

    padding-bottom: 10px;
    margin-bottom: 24px;
  }
}
h1 span {
  display: block;
}
.pink {
  color: #E22395;
  font-family: 'Monoton', sans-serif;
  margin-bottom: -110px;
  font-size: 148px;
  transform: translateY(-34px);
}

@media (min-width: 930px) {
  .pink{
    margin-bottom: -85px;
  }
}
.blue {
  color: #1AD6C7;
  text-transform: uppercase;
  font-family: 'Raleway', sans-serif;
  letter-spacing: 5px;
  position: relative;
  /*bottom: 60px;*/
  /*left: 50px;*/
  font-weight: bold;
  /*margin-bottom: -50px;*/

  transform: translateX(74px);
  font-size: 49px;
  display: inline-block;
}
@media (min-width: 930px) {
  .blue{
    font-size: 38px;
    /*left:auto;*/
    /*font-size: */
    /*right:0;*/
  }
}

h3{
  line-height: 19px;
  padding: 0 20px;
}
@media (min-width: 930px) {
  h3{
    padding: 0;
  }
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
.contribute.button {
  margin: 25px 20px 25px;
  display: block;
}

@media (min-width: 600px) {
  .contribute.button{
    display: inline-block;
    width: 175px;
    margin: 10px auto 20px 20px;
  }
}

@media (min-width: 930px) {
  .contribute.button {
    margin: 10px 0 25px;
  }
}

@media (min-width: 930px) {
  .about {
    width: 29%;
    margin-right: 0;
    padding-right: 20px;
    box-sizing: border-box;
  }
  /*h1 {
    margin-top: -60px;
  }*/
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

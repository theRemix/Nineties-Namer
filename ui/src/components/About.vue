<template>
  <div class="about">
    <div class="inner">
      <h1>
        <span class="pink">90s</span>
        <span class="blue">Namer</span>
      </h1>
      <h3>Generate random names from the 90s! <span>as a service™</span></h3>
      <p>Need some randomly generated names? 90s Namer will generate random names based on a hash or any input.</p>
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
  font-size: 50px;
  text-align: center;
  text-transform: uppercase;
  margin: 0;
}
h1 span {
  display: block;
}
.pink {
  color: #E22395;
}
.blue {
  color: #1AD6C7;
}
h3 span {
  font-style: italic;
}
h2, h3 {
  font-weight: normal;
}
.about {
  font-size: 14px;
}
.about .inner {
  padding-top: 0;
}
@media (min-width: 768px) {
  .about {
    width: 30%;
    margin-right: 20px;
  }
}
</style>

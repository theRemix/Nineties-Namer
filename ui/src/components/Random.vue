<template>
  <div class="random">
    <h2>Click to generate new random name.</h2>
    <div>
      <button type="button" v-on:click="generate">Generate</button>
    </div>
    <div class="output">
      <h3>{{randomName}}</h3>
    </div>
    <div class="errors">
      <p>{{errors}}</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'random',
  data: () => ({
    randomName: '',
    errors: ''
  }),
  methods: {
    generate: function (event) {
      fetch('/random')
        .then(res => {
          if (!res.ok) {
            throw res
          } else {
            return res.text()
          }
        })
        .then(randomName => {
          this.randomName = randomName
        })
        .catch(error => {
          this.errors = 'An error occurred when requesting the API.'
          console.error(error)
        })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h2, h3 {
  font-weight: normal;
}
</style>

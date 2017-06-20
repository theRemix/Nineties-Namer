<template>
  <div class="random">
    <h2>Generate a name</h2>
    <div class="inner">
      <div>
        <button type="button" v-on:click="generate" class='generate-button'>Generate</button>
      </div>
      <input placeholder="result" v-model="randomName" class="output" disabled />
      <div class="errors">
        <p>{{errors}}</p>
      </div>
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
.generate-button {
  border-color: #E02394;
  color: #E02394;
  background: transparent;
  font-size: 16px;
  margin-bottom: 30px;
}
.generate-button:hover {
  color: #FFFFFF;
  background: #E02394;
}
</style>

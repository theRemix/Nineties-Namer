<template>
  <div class="hash">
    <h2>Or enter some key to create a hashed random phrase</h2>
    <div>
      <input type="text" placeholder="your keywords (ex. v1.0.3)" v-on:keyup="generate" v-model="key" />
    </div>
    <div class="output">
      <h3>{{hashedName}}</h3>
    </div>
    <div class="errors">
      <p>{{errors}}</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'hash',
  data: () => ({
    key: '',
    hashedName: '',
    errors: ''
  }),
  methods: {
    generate: function (event) {
      if (this.key.length === 0) {
        this.hashedName = ''
        return
      }

      fetch(`/names/${this.key}`)
        .then(res => {
          if (!res.ok) {
            throw res
          } else {
            return res.text()
          }
        })
        .then(hashedName => {
          this.hashedName = hashedName
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

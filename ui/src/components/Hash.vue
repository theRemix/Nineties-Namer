<template>
  <div class="hash">
    <h2>Or enter some key <span class="show-for-large">to create a hashed random phrase</span></h2>
    <div class="inner">
      <div>
        <input type="text" placeholder="your keywords (ex. v1.0.3)" v-on:keyup="generate" v-model="key" />
      </div>
      <input class="output" placeholder="result" v-model="hashedName" disabled/>
      <div class="errors">
        <p>{{errors}}</p>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: 'hash',
  data: () => {
    return {
      key: window.location.hash !== '' ? decodeURI(window.location.hash.substr(1)) : '',
      hashedName: '',
      errors: ''
    }
  },
  created: function () {
    const loadName = key => fetch(`/names/${key}`)
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

    window.onhashchange = () => {
      this.key = decodeURI(window.location.hash.substr(1))
      if (this.key.length === 0) {
        this.hashedName = ''
        return
      }
      loadName(this.key)
    }

    if (this.key.length > 0) {
      loadName(this.key)
    }
  },
  methods: {
    generate: function () {
      window.location.hash = this.key
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h2, h3 {
  font-weight: normal;
}
.show-for-large {
  display: none;
}
@media(min-width: 930px) {
  .show-for-large {
    display: inline;
  }
}
</style>

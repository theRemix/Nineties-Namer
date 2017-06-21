<template>
  <div class="content-curl code">
    <span class="code-syntax-pink curl">curl 90s.fun/random</span>
    <span class="code-syntax-blue curl">curl 90s.fun/names/v1.2.0-rc8</span>
    <span class="curl">curl 90s.fun/names/2017-06-14T05:10:48.935Z</span>
  </div>
</template>

<script>
export default {
  name: 'content-curl'
}
</script>

<style scope>
  .curl {
    display: block;
    margin-bottom: 10px;
  }
</style>

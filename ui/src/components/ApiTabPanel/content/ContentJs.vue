<template>
  <div class="content-js">
    <code><pre>fetch('/random')
  .then( res => res.text() )
  .then( randomName => {
    console.log( randomName );
  })
</pre></code>

    <i>See: <a href="https://github.com/theRemix/Nineties-Namer/blob/master/ui/src/components/Random.vue">Random Example on Github</a></i>

    <code><pre>const someKey = 'v1.2.0-rc8';
fetch(`/names/${someKey}`)
  .then( res => res.text() )
  .then( generatedName => {
    console.log( generatedName );
  })
</pre></code>

    <i>See: <a href="https://github.com/theRemix/Nineties-Namer/blob/master/ui/src/components/Hash.vue">Names Example on Github</a></i>

  </div>
</template>

<script>
export default {
  name: 'content-js'
}
</script>


<template>
  <div class="content-js">
    <div class="code">
      <span><span class="code-syntax-blue">fetch</span>(<span class="code-syntax-pink"> '/random' </span>)</span>
      <div class="code-inner">
        <span>.<span class="code-syntax-blue">then </span>(<span class="code-syntax-pink"> res </span>=> <span class="code-syntax-pink">res</span>.text() )</span>
        <br>
        <span>.<span class="code-syntax-blue">then </span>(<span class="code-syntax-pink"> randomName</span> => {</span>
        <div class="code-inner">
          <span><span class="code-syntax-blue">console</span>.<span class="code-syntax-blue">log</span>( <span class="code-syntax-pink">randomName  </span>);</span>
        </div>
        <span>})</span>
      </div>
    </div>
    
    <br>
    <span><span class="code-syntax-blue">const someKey </span>= <span class="code-syntax-pink">'v1.2.0-rc8'</span>;</span>
    <div class="code">
      <span><span class="code-syntax-blue">fetch</span>(<span class="code-syntax-pink">`/names/</span>${<span class="code-syntax-blue">someKey</span>}<span class="code-syntax-pink">`</span>)</span>
      <br>
      <div class="code-inner">
        <span>.<span class="code-syntax-blue">then</span>( <span class="code-syntax-pink">res</span> => <span class="code-syntax-pink">res</span>.text() )</span>
        <br>
        <span>.<span class="code-syntax-blue">then</span>( <span class="code-syntax-pink">generatedName</span> => {</span>
        <div class="code-inner">
          <span><span class="code-syntax-blue">console</span>.<span class="code-syntax-blue">log</span>(<span class="code-syntax-pink"> generatedName</span> );</span>
        </div>
        <span>})</span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'content-js'
}
</script>


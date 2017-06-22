<template>
  <div class="content-js">
    <div class="code">
      <span><span class="code-syntax-blue">fetch</span>(<span class="code-syntax-pink"> &quot;https://90s.fun/random&quot; </span>)</span>
      <div>
        <span>&nbsp;&nbsp;.<span class="code-syntax-blue">then</span>(<span class="code-syntax-pink"> res </span>=> <span class="code-syntax-pink">res</span>.text() )</span>
        <br>
        <span>&nbsp;&nbsp;.<span class="code-syntax-blue">then</span>(<span class="code-syntax-pink"> randomName</span> => {</span>
        <div>
          <span>&nbsp;&nbsp;&nbsp;&nbsp;<span class="code-syntax-blue">console</span>.<span class="code-syntax-blue">log</span>( <span class="code-syntax-pink">randomName</span> )</span>
        </div>
        <span>&nbsp;&nbsp;})</span>
      </div>
    </div>

    <br>
    <span><span class="code-syntax-blue">const someKey </span>= <span class="code-syntax-pink">&quot;v1.2.0-rc8&quot;</span>;</span>
    <div class="code">
      <span><span class="code-syntax-blue">fetch</span>( <span class="code-syntax-pink">&#96;https://90s.fun/names/</span>${<span class="code-syntax-blue">someKey</span>}<span class="code-syntax-pink">&#96;</span> )</span>
      <br>
      <div>
        <span>&nbsp;&nbsp;.<span class="code-syntax-blue">then</span>( <span class="code-syntax-pink">res</span> => <span class="code-syntax-pink">res</span>.text() )</span>
        <br>
        <span>&nbsp;&nbsp;.<span class="code-syntax-blue">then</span>( <span class="code-syntax-pink">generatedName</span> => {</span>
        <div>
          <span>&nbsp;&nbsp;&nbsp;&nbsp;<span class="code-syntax-blue">console</span>.<span class="code-syntax-blue">log</span>(<span class="code-syntax-pink"> generatedName</span> )</span>
        </div>
        <span>&nbsp;&nbsp;})</span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'content-js'
}
</script>


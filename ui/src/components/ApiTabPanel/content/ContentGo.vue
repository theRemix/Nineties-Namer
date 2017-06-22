<template>
  <div class="content-go">
    <div class="code">
      <span class="code-syntax-blue">import </span><span>(</span>
      <div>
        <span class="code-syntax-pink">&nbsp;&nbsp;&quot;fmt&quot;</span>
        <br>
        <span class="code-syntax-pink">&nbsp;&nbsp;&quot;github.com/theRemix/go90s&quot;</span>
      </div>
      <span>)</span>
    </div>
    <div class="code">
      <span class="code-syntax-blue">func </span><span class="code-syntax-pink">main </span><span>() {</span>
      <div>
        <span>&nbsp;&nbsp;fmt.Println(<span class="code-syntax-blue">go90s.GetRandomName</span>())</span>
        <br>

        <span>&nbsp;&nbsp;fmt.Println(<span class="code-syntax-blue">go90s.GetName</span>(<span class="code-syntax-pink">&quot;v1.2.0-rc8&quot;</span>))</span>
      </div>
     <span>}</span>
    </div>
  </div>
</template>

<script>
export default {
  name: 'content-go'
}
</script>

<template>
  <div class="content-go">
    <div class="code">
      <span class="code-syntax-blue">import </span><span>(</span>
      <div class="code-inner">
        <span class="code-syntax-pink">fmt"</span>
        <br>
        <span class="code-syntax-pink">"github.com/theRemix/go90s"</span>
      </div>
      <span>)</span>
    </div>
    <div class="code">
      <span class="code-syntax-blue">func </span><span class="code-syntax-pink">main </span><span>() {</span>
      <div class="code-inner">
        <span>fmt.Println(<span class="code-syntax-blue">go90s.GetRandomName></span>())</span>
        <br>

        <span>fmt.Println(<span class="code-syntax-blue">go90s.GetName></span>(<span class="code-syntax-pink">v1.2.0-rc8"</span>))</span>
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

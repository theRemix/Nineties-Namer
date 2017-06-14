<template>
  <div class="content-go">
    <code><pre>import (
  "fmt"
  "github.com/theRemix/go90s"
)

func main() {
  fmt.Println(go90s.GetRandomName())

  fmt.Println(go90s.GetName("v1.2.0-rc8"))
}
    </pre></code>

    <i>See: <a href="https://github.com/theRemix/Nineties-Namer/blob/master/routes/router.go">Example on Github</a></i>

  </div>
</template>

<script>
export default {
  name: 'content-go'
}
</script>

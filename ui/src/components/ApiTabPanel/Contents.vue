<template>
  <div class="tab-contents">
    <content-curl v-if='activeTab === "curl"'></content-curl>
    <content-go v-if='activeTab === "go"'></content-go>
    <content-js v-if='activeTab === "js"'></content-js>
  </div>
</template>

<script>
import ContentCurl from './content/ContentCurl'
import ContentGo from './content/ContentGo'
import ContentJs from './content/ContentJs'

export default {
  name: 'contents',
  props: ['active-tab'],
  components: {
    ContentCurl,
    ContentGo,
    ContentJs
  }
}
</script>

<style>
.tab-contents {
  font-size: 16px;
  font-weight: bold;
  font-family: 'Inconsolata', monospace;
}
.code {
  margin-bottom: 10px;
}
.code-inner {
  padding-left: 10px;
}
.code-syntax-blue {
  color: #1AD6C7;
}
.code-syntax-pink {
  color: #E22395;
}
i {
  color: #000000;
  font-style: normal;
}
i a {
  color: #E22395;
}
</style>


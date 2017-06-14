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
code{

}
pre{

}
</style>


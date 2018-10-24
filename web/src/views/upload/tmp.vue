<template>
  <div>
    <h1>{{ msg }}</h1>
  </div>
</template>

<script>
export default {
  name: 'test',
  data() {
    return {
      msg: 'tmp'
    }
  },
  created() {
    this.get_login_info()
  },
  methods: {
    get_login_info() {
    }
  }
}
</script>

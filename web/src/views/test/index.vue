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
      msg: 'test index'
    }
  },
  created() {
    this.get_login_info()
  },
  methods: {
    get_login_info() {
      this.axios.get('http://localhost:8080/Login/GetUserInfo')
        .then(function(response) {
          console.log(response)
          this.msg = response
        })
        .catch(function(error) {
          console.log(error)
        })
    }
  }
}
</script>

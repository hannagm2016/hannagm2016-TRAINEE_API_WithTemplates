<template>
<div>

<main class="form-signin">
    <form role="form" method="POST" action="http://localhost:8000/authorisationPost">
        <h1 class="h3 mb-3 fw-normal">Please sign in</h1>
        <!--<label for="inputEmail" class="visually-hidden">Email address</label>-->
        <input type="email" id="inputEmail"  name="inputEmail" class="form-control" placeholder="Email address" required="" autofocus="">
       <!-- <label for="inputPassword" class="visually-hidden">Password</label>-->
        <input type="password" id="inputPassword" name="inputPassword" class="form-control" placeholder="Password" required="">
        <div class="checkbox mb-3">
            <label>
                <input type="checkbox" value="remember-me"> Remember me
            </label>
        </div>
        <button class="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
    </form>

    <title>Golang Login Facebook/Google</title>
<p>{{link.FB}}</p>
        <button>Login with Facebook!</button> </a>
    <p>{{link.Google}}</p>
        <button>Login with Google!</button> </a>
    <div class="text-muted py-2" style="text-align:center;">

</div>

    <div class="text-muted py-3">
        <a href="/#/registration">Registration</a>
    </div>
</main>
<hr>
  </div>
</template>

<script>
export default {
  data: function(){
    return {
      link: null
    }
  },
  watch: {
    '$route' (to,from){
      this.getData()
    }
  },
  methods: {
    getData (){
     var that=this
      axios.get ('http://localhost:8000/authorization')
        .then(function (response) {
        that.link = response.data
        console.log (link)
        })
    }
  },
  mounted: function (){
    this.getData()
  }
}
</script>



<style>

.form-signin {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: auto;
  margin-top:100px;
 }
.form-signin .checkbox {
  font-weight: 400;
}
.form-signin .form-control {
  position: relative;
  box-sizing: border-box;
  height: auto;
  padding: 10px;
  font-size: 16px;
}
.form-signin .form-control:focus {
  z-index: 2;
}
.form-signin input[type="email"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}
.form-signin input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}

</style>

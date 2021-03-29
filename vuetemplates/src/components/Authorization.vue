<template>
<div>

<main class="form-signin">
 <form @submit.prevent="authorizeUser">
         <h1 class="h3 mb-3 fw-normal">Please Registrate</h1>
          <input v-model="User.Email" type="email" id="inputEmail"  name="inputEmail" class="form-control" placeholder="Email address" required="" autofocus="">
          <input v-model="User.Password" type="password" id="inputPassword" name="inputPassword" class="form-control" placeholder="Password" required="">
         <button class="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
          <label>
               <input type="checkbox" value="remember-me"> Remember me
          </label>
                 </form>
       <div class="text-muted py-2">
        <a v-bind:href="link.FB">
        <button>Login with Facebook!</button> </a>
         <a v-bind:href= "link.Google">
        <button>Login with Google!</button> </a>
</div>

    <div class="text-muted py-3">
        <a href="/#/registration">Registration</a>
    </div>
</main>
<hr>
  </div>
</template>

<script>
import axios from 'axios'
export default {
   data: () => ({
      User: {
        Email: '',
        Password: '123',
      },
       link: []
    }),
    mounted() {
       fetch("http://localhost:8000/authorization")
         .then(response =>response.json())
         .then((data)=>{
          this.link = data;
         })
     },
      methods: {
          authorizeUser() {
           axios.post (`http://localhost:8000/authorizationPost`, {
                          Email: this.User.Email,
                          Password: this.User.Password,

                     })
                      .then (response => {
                     console.log (response)
                      alert ("Authorized: " + this.User.Email)

                      })
                       window.location = '/'
                  }
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

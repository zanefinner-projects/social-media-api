const forms = {
    loginSignupForm: 
        `
        <form method="post">
        Username: <input name = "username"/>
        Password: <input name = "password"/>
        `
}


var app = new Vue({
    el: "#app",
    data: {
        message: forms.loginSignupForm,
        location: "Log in or Sign Up"
    }
})

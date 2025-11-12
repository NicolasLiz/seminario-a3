document.getElementById("loginButton").addEventListener("click", onSubmit)

async function getPublicKey() {
    const res = await fetch("/api/publicKey");
    const data = await res.json();
    const publicKey = data.publicKey

    return publicKey
}

async function onSubmit() {
    const publicKey = await getPublicKey();
    const form = document.getElementById("loginForm")
    const formData = {
        email: form.elements["email"].value,
        senha: form.elements["senha"].value
    };

    const encrypt = new JSEncrypt();
    encrypt.setPublicKey(publicKey);
    const encrypted = encrypt.encrypt(JSON.stringify(formData));

    await fetch("/submit", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({encrypted})
    }).then(response => response.text())
        .then(url => {
            // Redirect the browser manually
            window.location.href = url; 
        })
}

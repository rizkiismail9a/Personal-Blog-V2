function sendEmail() {
  let name = document.getElementById("Name").value;
  let email = document.getElementById("Email").value;
  let phoneNumber = document.getElementById("Phone").value;
  let subject = document.getElementById("Subject").value;
  let message = document.getElementById("Message").value;

  if (name == "") {
    alert("Isi namanya, Boi.");
  } else if (email == "") {
    alert("Isi alamat emailnya, Boi.");
  } else if (phoneNumber == "") {
    alert("Isi nomor hp-nya, Boi.");
  } else if (subject == "") {
    alert("Mau ngomongin apa?");
  } else if (message == "") {
    alert("Mau ngomong apa?");
  } else {
    let penerima = "muhamadrizkiismail9a@gmail.com";
    let a = document.createElement("a");
    a.href = `mailto:${penerima}?subject=${subject}&body=Halo, nama saya ${name}. ${message}. Jika Anda ingin melanjutkan percakapan, tolong hubungi saya di nomor ${phoneNumber} atau ${email}. Terima kasih`;
    a.click();
  }
}

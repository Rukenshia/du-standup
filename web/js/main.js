const parts = window.location.href.split("/");

window.baseURL = `${parts[0]}//${parts[2]}`;

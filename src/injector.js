const site = window.location.hostname

function injectJSCode(code) {
	const scriptElement = document.createElement('script');
	scriptElement.setAttribute('type', 'text/javascript');
	scriptElement.textContent = code;
	document.documentElement.appendChild(scriptElement);
}
function injectJSLink(src) {
	const scriptElement = document.createElement('script');
	scriptElement.setAttribute('type', 'text/javascript');
	scriptElement.setAttribute('src', src);
	document.documentElement.appendChild(scriptElement);
}

if (site == 'https://emat.dk/') {
	const custom_element = document.createElement('h1')
    custom_element.innerHTML = 'value'
    document.body.append(custom_element)
}
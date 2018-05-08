// document.addEventListener('DOMContentLoaded', function() {
    document.getElementById("btn-close-app").addEventListener("click", goToMain);
// });

function goToMain() {
    location.href = "/";
};

//https://stackoverflow.com/questions/49244944/toggle-appendchild-if-doesnt-exist-and-removechild-if-exist/49245337#49245337
//Will play with this to use display: block and outside click to hide.
let element = document.getElementById('os-start-menu');
let content = element.content;

function toggle() {
    'use strict';
    // Attempt to reference the element in the document, not the template content
    var imported = document.querySelector(".os1-start-menu");

    // Check for the element, not the template content
    if (document.body.contains(imported)) {
        // Element exists, get its parent
        imported.parentNode.removeChild(imported);
    } else {
        // Use .importNode to bring template content in:
        document.body.appendChild(document.importNode(content, true));
    }
}

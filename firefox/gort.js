async function getCurrentUrl() {
  let currentUrl;
  await browser.tabs
    .query({ currentWindow: true, active: true })
    .then((tabs) => {
      let tab = tabs[0];
      currentUrl = tab.url;
    }, console.error);
  return currentUrl;
}

const shorten = async () => {
  let currentUrl = await getCurrentUrl();
  try {
    const response = await fetch("http://127.0.0.1:6969", {
      method: "POST",
      headers: {
        "Content-Type": "text/plain",
      },
      body: JSON.stringify({ long_url: currentUrl }),
    });

    const result = await response.json();
    console.log("Success:", result);
  } catch (error) {
    console.error("Error:", error);
  }
};

browser.browserAction.onClicked.addListener(shorten);

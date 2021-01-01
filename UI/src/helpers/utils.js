export function getStorageItem(name = 'user') {
    const item = localStorage.getItem(name);
    if (item !== 'undefined') {
        return JSON.parse(item) || {};
    } else {
        return {};
    }
}

export function setStorageItem(name, data) {
    localStorage.setItem(name, JSON.stringify(data));
}

export function removeStorageItem(name) {
    localStorage.removeItem(name);
}

export function clearStorage() {
    localStorage.clear();
}


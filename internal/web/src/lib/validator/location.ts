import { LngLat } from '../../types/data-types';

export function parseLocation(str: string): [LngLat, boolean] {
    // Regular expression to match the format "lat,lng"
    const regex = /^(-?\d+(\.\d+)?),\s*(-?\d+(\.\d+)?)$/;

    const match = str.match(regex);

    if (match) {
        const lat = parseFloat(match[1]);
        const lng = parseFloat(match[3]);

        // Check if the parsed values are valid
        if (
            isNaN(lat) ||
            isNaN(lng) ||
            lat < -90 ||
            lat > 90 ||
            lng < -180 ||
            lng > 180
        ) {
            return [{ lat: 0, lng: 0 }, false];
        }

        return [{ lat, lng }, true];
    }

    // If the input doesn't match the expected format
    return [{ lat: 0, lng: 0 }, false];
}

export function isLocation(str: string): boolean {
    // Regular expression to match the format "lat,lng"
    const regex = /^(-?\d+(\.\d+)?),\s*(-?\d+(\.\d+)?)$/;
    const match = str.match(regex);

    if (match) {
        const lat = parseFloat(match[1]);
        const lng = parseFloat(match[3]);

        // Check if the parsed values are valid
        if (
            isNaN(lat) ||
            isNaN(lng) ||
            lat < -90 ||
            lat > 90 ||
            lng < -180 ||
            lng > 180
        ) {
            return false;
        }

        return true;
    }

    return false;
}

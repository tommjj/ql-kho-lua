'use client';
import Map, { MapRef, NavigationControl, Popup } from 'react-map-gl/maplibre';
import {} from 'react-map-gl/maplibre';
import 'maplibre-gl/dist/maplibre-gl.css';
import '@maplibre/maplibre-gl-geocoder/dist/maplibre-gl-geocoder.css';
import { useEffect, useRef } from 'react';
import MaplibreGeocoder from '@maplibre/maplibre-gl-geocoder';
import maplibregl from 'maplibre-gl';
import { geocoder } from '@/lib/map/geocoder';

const API_KEY = process.env.NEXT_PUBLIC_MAP_STYLE_API_KEY;

function MapContainer() {
    const mapRef = useRef<MapRef | undefined>(undefined);

    useEffect(() => {}, []);

    return (
        <div className="relative w-full h-screen ">
            <Map
                onLoad={() => {
                    const map = mapRef.current;
                    if (!map) {
                        return;
                    }

                    const geo = new MaplibreGeocoder(geocoder, {
                        maplibregl: maplibregl,
                        zoom: 14,
                    });
                    map.addControl(geo, 'top-right');
                }}
                // eslint-disable-next-line @typescript-eslint/no-explicit-any
                ref={mapRef as any}
                initialViewState={{
                    longitude: -100,
                    latitude: 40,
                    zoom: 3.5,
                }}
                mapStyle={`https://api.maptiler.com/maps/streets-v2/style.json?key=${API_KEY}`}
            >
                <NavigationControl position="top-left" />
                <Popup longitude={-100} latitude={40} anchor="bottom">
                    You are here
                </Popup>
            </Map>
        </div>
    );
}

export default MapContainer;

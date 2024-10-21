'use client';

import Map, { Marker } from 'react-map-gl/maplibre';
import Pin from '@/components/map/pin';
import { Warehouse } from '@/lib/zod.schema';
import { cn } from '@/lib/utils';

import 'maplibre-gl/dist/maplibre-gl.css';

const API_KEY = process.env.NEXT_PUBLIC_MAP_STYLE_API_KEY;

function WarehouseMap({
    warehouse,
    className,
}: {
    warehouse: Warehouse;
    className?: string;
}) {
    return (
        <div className={cn('relative w-full h-full', className)}>
            <Map
                initialViewState={{
                    zoom: 5,
                    latitude: warehouse.location[0],
                    longitude: warehouse.location[1],
                }}
                mapStyle={`https://api.maptiler.com/maps/streets-v2/style.json?key=${API_KEY}`}
            >
                <Marker
                    longitude={warehouse.location[1]}
                    latitude={warehouse.location[0]}
                    anchor="bottom"
                >
                    <Pin />
                </Marker>
            </Map>
        </div>
    );
}

export default WarehouseMap;

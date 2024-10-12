'use client';

import { Storehouse } from '@/lib/zod.schema';
import {
    ResizableHandle,
    ResizablePanel,
    ResizablePanelGroup,
} from '../../shadcn-ui/resizable';
import MapContainer, { MapRefType } from '../../map/map';
import { useRef, useState } from 'react';
import { Marker, Popup } from 'react-map-gl/maplibre';
import Pin from '../../map/pin';
import StorehousePopup from '../../map/strorehouse';

import StorehouseList from './storehouse-list';

type Props = {
    stores: Storehouse[];
};

function StorehousePage({ stores }: Props) {
    const mapRef = useRef<MapRefType>();
    const [storeInfo, setStoreInfo] = useState<Storehouse | null>(null);

    return (
        <section className="flex w-full h-screen">
            <ResizablePanelGroup
                direction="horizontal"
                className="w-full rounded-none md:min-w-[450px]"
            >
                <ResizablePanel
                    className="relative min-w-[220px] "
                    defaultSize={80}
                    minSize={50}
                    maxSize={80}
                >
                    <MapContainer
                        // eslint-disable-next-line @typescript-eslint/no-explicit-any
                        ref={mapRef as any}
                    >
                        {stores.map((store) =>
                            store !== storeInfo ? (
                                <Marker
                                    key={store.id}
                                    longitude={store.location[1]}
                                    latitude={store.location[0]}
                                    anchor="bottom"
                                    onClick={(e) => {
                                        e.originalEvent.stopPropagation();
                                        setStoreInfo(store);
                                    }}
                                >
                                    <Pin />
                                </Marker>
                            ) : null
                        )}

                        {storeInfo ? (
                            <Popup
                                longitude={storeInfo.location[1]}
                                latitude={storeInfo.location[0]}
                                anchor="bottom"
                                onClose={() => setStoreInfo(null)}
                            >
                                <StorehousePopup storehouse={storeInfo} />
                            </Popup>
                        ) : null}
                    </MapContainer>
                </ResizablePanel>
                <ResizableHandle withHandle />
                <ResizablePanel defaultSize={20}>
                    <StorehouseList
                        storehouses={stores}
                        mapLocationControl={(longitude, latitude) => {
                            const map = mapRef.current;
                            if (!map) return;

                            map.flyTo({
                                center: [longitude, latitude],
                                zoom: 14,
                            });
                        }}
                    />
                </ResizablePanel>
            </ResizablePanelGroup>
        </section>
    );
}

export default StorehousePage;

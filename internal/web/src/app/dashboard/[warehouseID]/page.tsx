import { authz } from '@/auth';
import { CapacityPieChart } from '@/components/pages/warehouse/chart';
import InventoryList from '@/components/pages/warehouse/inventory-list';
import WarehouseMap from '@/components/pages/warehouse/warehouse-map';
import Header from '@/components/ui/header';
import SearchBar from '@/components/ui/search';
import { UpdateWarehouse } from '@/components/warehouse/update-warehouse';
import { ErrDataNotFound, ErrUnauthorized } from '@/lib/errors';
import { handleErr } from '@/lib/response';
import {
    getWarehouseInventory,
    getWarehouseByID,
} from '@/lib/services/warehouse.service';
import { Role } from '@/types/role';

type Props = {
    params: {
        warehouseID: string;
    };
};

async function StorePage({ params: { warehouseID } }: Props) {
    const user = await authz();
    if (!user) {
        handleErr(ErrUnauthorized);
    }

    let id = 0;
    if (warehouseID) {
        id = Number(warehouseID);
        if (!Number.isInteger(id)) {
            handleErr(ErrDataNotFound);
        }
    }

    const [warehouseRes, err] = await getWarehouseByID(user.token, id);
    if (err) {
        handleErr(err);
    }
    const warehouse = warehouseRes.data;

    const [invRes, err2] = await getWarehouseInventory(user.token, id);
    if (err2) {
        handleErr(err2);
    }
    const inv = invRes.data;

    return (
        <section className="relative w-full h-screen">
            <Header
                className="absolute left-0 top-0 w-full"
                title={warehouse.name}
            >
                {user.role === Role.ROOT ? (
                    <UpdateWarehouse warehouse={warehouse} />
                ) : null}
            </Header>
            <div className="flex w-full h-full">
                <div className="flex flex-col flex-grow pt-[3.8rem] px-2">
                    <SearchBar shallow />
                    <div className="flex-grow pt-1">
                        <InventoryList warehouseItems={inv} />
                    </div>
                </div>

                <div className="flex flex-col h-full p-2 gap-2 pt-[3.8rem] w-80">
                    <div className="">
                        <CapacityPieChart
                            warehouse={warehouse}
                            warehouseItems={inv}
                        />
                    </div>
                    <div className="flex-grow h-auto overflow-hidden rounded-xl shadow border">
                        <WarehouseMap warehouse={warehouse} />
                    </div>
                </div>
            </div>
        </section>
    );
}

export default StorePage;

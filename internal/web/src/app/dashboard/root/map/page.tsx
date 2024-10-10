import { authz } from '@/auth';
import MapContainer from '@/components/map/map_test';
import NoSSR from '@/components/noSSR';
import { ErrUnauthorized } from '@/lib/errors';
import fetcher from '@/lib/http/fetcher';
import { handleErr } from '@/lib/response';

async function MapPage() {
    const user = await authz();
    if (!user) {
        handleErr(ErrUnauthorized);
    }

    const [data, err] = await fetcher
        .set('authorization', `jwt ${user.token}`)
        .get('/v1/api/storehouses?limit=999');

    if (err) {
        handleErr(err);
    }

    console.log(data);

    return (
        <section>
            <NoSSR>
                <MapContainer />
            </NoSSR>
        </section>
    );
}
export default MapPage;

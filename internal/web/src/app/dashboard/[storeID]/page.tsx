type Props = {
    params: {
        storeID: string;
    };
};

function StorePage({ params: { storeID } }: Props) {
    return (
        <div>
            <h1>{storeID}</h1>
        </div>
    );
}

export default StorePage;

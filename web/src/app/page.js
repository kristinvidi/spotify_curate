'use client';

import { useState, useEffect } from 'react';

export default function Home() {
    const [userID, setUserID] = useState(null);
    const [loading, setLoading] = useState(true);
    const [syncing, setSyncing] = useState(false);
    const [tags, setTags] = useState([]);
    const [artists, setArtists] = useState([]);
    const [selectedTag, setSelectedTag] = useState('');
    const [checkedArtists, setCheckedArtists] = useState(new Set());
    const [toast, setToast] = useState(null);

    // Initialize
    useEffect(() => {
        async function init() {
            try {
                const res = await fetch('/api/init');
                const data = await res.json();

                if (data.userID) {
                    setUserID(data.userID);
                    await fetchData(data.userID);
                } else {
                    showToast('Failed to authenticate. Please check backend.', 'error');
                }
            } catch (err) {
                showToast('Connection failed: ' + err.message, 'error');
            } finally {
                setLoading(false);
            }
        }
        init();
    }, []);

    async function fetchData(uid) {
        try {
            const [tagsRes, artistsRes] = await Promise.all([
                fetch(`/api/tags?userID=${uid}`),
                fetch(`/api/artists/unmapped?userID=${uid}`)
            ]);

            const tagsData = await tagsRes.json();
            const artistsData = await artistsRes.json();

            if (tagsData.tags) setTags(tagsData.tags);
            if (artistsData.artists) setArtists(artistsData.artists);
        } catch (err) {
            console.error(err);
            showToast('Failed to load data', 'error');
        }
    }

    const handleSync = async () => {
        setSyncing(true);
        try {
            const res = await fetch('/api/sync', { method: 'POST' });
            const data = await res.json();
            if (data.success) {
                showToast('Data synced successfully', 'success');
                await fetchData(userID);
            } else {
                throw new Error(data.error);
            }
        } catch (err) {
            showToast('Sync failed: ' + err.message, 'error');
        } finally {
            setSyncing(false);
        }
    };

    const handleToggleArtist = (artistName) => {
        const newSet = new Set(checkedArtists);
        if (newSet.has(artistName)) {
            newSet.delete(artistName);
        } else {
            newSet.add(artistName);
        }
        setCheckedArtists(newSet);
    };

    const handleMap = async () => {
        if (!selectedTag || checkedArtists.size === 0) {
            showToast('Select a tag and at least one artist', 'error');
            return;
        }

        try {
            const res = await fetch('/api/mappings', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    userID,
                    tag: selectedTag,
                    artists: Array.from(checkedArtists),
                }),
            });
            const data = await res.json();

            if (data.success) {
                showToast(`Mapped ${checkedArtists.size} artists to ${selectedTag}`, 'success');
                // Remove mapped artists from list
                setArtists(artists.filter(a => !checkedArtists.has(a.name)));
                setCheckedArtists(new Set());
            } else {
                throw new Error(data.error);
            }
        } catch (err) {
            showToast('Mapping failed: ' + err.message, 'error');
        }
    };

    const handleCreatePlaylist = async () => {
        if (!selectedTag) {
            showToast('Select a tag first', 'error');
            return;
        }

        try {
            showToast('Creating playlist...', 'success');
            const res = await fetch('/api/playlists', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ tag: selectedTag }),
            });
            const data = await res.json();

            if (data.success) {
                showToast(`Playlist for ${selectedTag} created!`, 'success');
            } else {
                throw new Error(data.error);
            }
        } catch (err) {
            showToast('Playlist creation failed: ' + err.message, 'error');
        }
    };

    const showToast = (message, type) => {
        setToast({ message, type });
        setTimeout(() => setToast(null), 3000);
    };

    if (loading) return <div className="container"><h2>Connecting...</h2></div>;

    return (
        <div className="container">
            <header>
                <h1>Spotify Curator</h1>
                <div className="controls">
                    <button className="btn btn-secondary" onClick={handleSync} disabled={syncing}>
                        {syncing ? 'Syncing...' : 'Sync Data'}
                    </button>
                </div>
            </header>

            {userID ? (
                <div className="dashboard">
                    {/* Left Column: Unmapped Artists */}
                    <div className="card">
                        <h2>Unmapped Artists ({artists.length})</h2>
                        <div className="artist-list">
                            {artists.length === 0 ? (
                                <p style={{ padding: '1rem', color: '#888' }}>All artists mapped!</p>
                            ) : (
                                artists.map((artist) => (
                                    <div key={artist.id} className="artist-item">
                                        <input
                                            type="checkbox"
                                            checked={checkedArtists.has(artist.name)}
                                            onChange={() => handleToggleArtist(artist.name)}
                                        />
                                        <span>{artist.name}</span>
                                    </div>
                                ))
                            )}
                        </div>
                    </div>

                    {/* Right Column: Actions */}
                    <div className="card">
                        <h2>Actions</h2>

                        <div className="form-group">
                            <label>Select Tag (Genre)</label>
                            <select
                                value={selectedTag}
                                onChange={(e) => setSelectedTag(e.target.value)}
                            >
                                <option value="">-- Select Tag --</option>
                                {tags.map((tag) => (
                                    <option key={tag} value={tag}>{tag}</option>
                                ))}
                            </select>
                        </div>

                        <div className="form-group">
                            <button
                                className="btn btn-primary"
                                onClick={handleMap}
                                disabled={!selectedTag || checkedArtists.size === 0}
                                style={{ width: '100%' }}
                            >
                                Map Selected Artists to Tag
                            </button>
                        </div>

                        <hr style={{ borderColor: 'var(--border)', margin: '2rem 0' }} />

                        <div className="form-group">
                            <label>Create Playlist</label>
                            <p style={{ fontSize: '0.8rem', color: '#888', marginBottom: '1rem' }}>
                                Generates a "Recent in {selectedTag || '...'}" playlist.
                            </p>
                            <button
                                className="btn btn-secondary"
                                onClick={handleCreatePlaylist}
                                disabled={!selectedTag}
                                style={{ width: '100%' }}
                            >
                                Create Playlist
                            </button>
                        </div>
                    </div>
                </div>
            ) : (
                <div className="card">
                    <h2>Authentication Required</h2>
                    <p>Could not connect to backend. Ensure it is running.</p>
                </div>
            )}

            {toast && (
                <div className="toast" style={{ borderColor: toast.type === 'error' ? 'var(--error)' : 'var(--success)' }}>
                    {toast.message}
                </div>
            )}
        </div>
    );
}
